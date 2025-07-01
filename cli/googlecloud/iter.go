package googlecloud

import (
	"errors"
	"fmt"
	"iter"

	"google.golang.org/api/iterator"
)

//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=mock_$GOFILE

type Iterator[T any] interface {
	Next() (T, error)
}

type APIResult[T any] struct {
	Item  T
	Error error
}

func buildIterSeq[APIItem any](
	apiIterator Iterator[APIItem],
) iter.Seq[APIResult[APIItem]] {
	return BuildConvertingIterSeq(
		apiIterator,
		func(item APIItem) APIItem { return item },
	)
}

func BuildConvertingIterSeq[APIItem, DDItem any](
	apiIterator Iterator[APIItem],
	ddConverter func(APIItem) DDItem,
) iter.Seq[APIResult[DDItem]] {
	return func(yield func(result APIResult[DDItem]) bool) {
		for {
			apiItem, err := apiIterator.Next()

			if err != nil {
				if errors.Is(err, iterator.Done) {
					return
				}

				_ = yield(APIResult[DDItem]{Error: fmt.Errorf("querying google api: %w", err)})
				return
			}

			if !yield(APIResult[DDItem]{Item: ddConverter(apiItem)}) {
				return
			}
		}
	}
}
