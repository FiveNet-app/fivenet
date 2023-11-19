// Modified version of https://github.com/HeavyHorst/certmagic-nats/blob/b27fd6c010166e396b6f9e1c651ba7b02ce6c01f/nats.go#L114
// which is licensed under [MIT License](https://github.com/HeavyHorst/certmagic-nats/blob/b27fd6c010166e396b6f9e1c651ba7b02ce6c01f/LICENSE)
package locks

import (
	"context"
	"encoding/binary"
	"errors"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

const keyPrefix = "LOCK."

type Locks struct {
	logger *zap.Logger
	kv     nats.KeyValue

	bucket string

	revMap  map[string]uint64
	maplock sync.Mutex
}

func New(logger *zap.Logger, kv nats.KeyValue, bucket string) (*Locks, error) {
	l := &Locks{
		logger: logger.Named("locks").With(zap.String("bucket", bucket)),
		kv:     kv,

		bucket: bucket,

		revMap:  map[string]uint64{},
		maplock: sync.Mutex{},
	}

	return l, nil
}

// Lock acquires the lock for key, blocking until the lock
// can be obtained or an error is returned. Note that, even
// after acquiring a lock, an idempotent operation may have
// already been performed by another process that acquired
// the lock before - so always check to make sure idempotent
// operations still need to be performed after acquiring the
// lock.
//
// The actual implementation of obtaining of a lock must be
// an atomic operation so that multiple Lock calls at the
// same time always results in only one caller receiving the
// lock at any given time.
//
// To prevent deadlocks, all implementations (where this concern
// is relevant) should put a reasonable expiration on the lock in
// case Unlock is unable to be called due to some sort of network
// failure or system crash.
func (l *Locks) Lock(ctx context.Context, key string) error {
	l.logger.Info("lock", zap.String("key", key))
	lockKey := keyPrefix + key

loop:
	for {
		// Check for existing lock
		revision, err := l.kv.Get(lockKey)
		if err != nil && !errors.Is(err, nats.ErrKeyNotFound) {
			return err
		}

		if revision == nil {
			break
		}

		expires := time.Unix(0, int64(binary.LittleEndian.Uint64(revision.Value())))
		// Lock exists, check if expired
		if time.Now().After(expires) {
			// the lock expired and can be deleted
			// break and try to create a new one
			l.setRev(lockKey, revision.Revision())
			if err := l.Unlock(ctx, key); err != nil {
				if isWrongSequence(err) {
					goto loop
				}
				return err
			}
			break
		}

		select {
		// retry after a short period of time
		case <-time.After(time.Duration(50+rand.Float64()*(200-50+1)) * time.Millisecond):
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	// lock doesn't exist, create it
	contents := make([]byte, 8)
	binary.LittleEndian.PutUint64(contents, uint64(time.Now().Add(time.Duration(5*time.Minute)).UnixNano()))
	nrev, err := l.kv.Create(lockKey, contents)
	if err != nil && isWrongSequence(err) {
		// another process created the lock in the meantime
		// try again
		goto loop
	}

	if err != nil {
		return err
	}

	l.setRev(lockKey, nrev)
	return nil
}

func (l *Locks) IsLocked(ctx context.Context, key string) (bool, error) {
	// Check for existing lock
	revision, err := l.kv.Get(key)
	if err != nil && !errors.Is(err, nats.ErrKeyNotFound) {
		return false, err
	}

	return revision == nil, nil
}

// Unlock releases the lock for key. This method must ONLY be
// called after a successful call to Lock, and only after the
// critical section is finished, even if it errored or timed
// out. Unlock cleans up any resources allocated during Lock.
func (l *Locks) Unlock(ctx context.Context, key string) error {
	l.logger.Info("unlock", zap.String("key", key))
	lockKey := keyPrefix + key
	return l.kv.Delete(lockKey, nats.LastRevision(l.getRev(lockKey)))
}

func (l *Locks) setRev(key string, value uint64) {
	l.maplock.Lock()
	defer l.maplock.Unlock()
	l.revMap[key] = value
}

func (l *Locks) getRev(key string) uint64 {
	l.maplock.Lock()
	defer l.maplock.Unlock()
	return l.revMap[key]
}

func isWrongSequence(err error) bool {
	return strings.Contains(err.Error(), "wrong last sequence")
}
