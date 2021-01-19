// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package datastoretest

import (
	"context"
	"github.com/cadmiumcat/books-api/config"
	"github.com/cadmiumcat/books-api/interfaces"
	"github.com/cadmiumcat/books-api/models"
	"sync"
)

// Ensure, that DataStoreMock does implement interfaces.DataStore.
// If this is not the case, regenerate this file with moq.
var _ interfaces.DataStore = &DataStoreMock{}

// DataStoreMock is a mock implementation of interfaces.DataStore.
//
//     func TestSomethingThatUsesDataStore(t *testing.T) {
//
//         // make and configure a mocked interfaces.DataStore
//         mockedDataStore := &DataStoreMock{
//             AddBookFunc: func(book *models.Book)  {
// 	               panic("mock out the AddBook method")
//             },
//             CloseFunc: func(ctx context.Context) error {
// 	               panic("mock out the Close method")
//             },
//             InitFunc: func(in1 config.MongoConfig) error {
// 	               panic("mock out the Init method")
//             },
//         }
//
//         // use mockedDataStore in code that requires interfaces.DataStore
//         // and then make assertions.
//
//     }
type DataStoreMock struct {
	// AddBookFunc mocks the AddBook method.
	AddBookFunc func(book *models.Book)

	// CloseFunc mocks the Close method.
	CloseFunc func(ctx context.Context) error

	// InitFunc mocks the Init method.
	InitFunc func(in1 config.MongoConfig) error

	// calls tracks calls to the methods.
	calls struct {
		// AddBook holds details about calls to the AddBook method.
		AddBook []struct {
			// Book is the book argument value.
			Book *models.Book
		}
		// Close holds details about calls to the Close method.
		Close []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// Init holds details about calls to the Init method.
		Init []struct {
			// In1 is the in1 argument value.
			In1 config.MongoConfig
		}
	}
	lockAddBook sync.RWMutex
	lockClose   sync.RWMutex
	lockInit    sync.RWMutex
}

// AddBook calls AddBookFunc.
func (mock *DataStoreMock) AddBook(book *models.Book) {
	if mock.AddBookFunc == nil {
		panic("DataStoreMock.AddBookFunc: method is nil but DataStore.AddBook was just called")
	}
	callInfo := struct {
		Book *models.Book
	}{
		Book: book,
	}
	mock.lockAddBook.Lock()
	mock.calls.AddBook = append(mock.calls.AddBook, callInfo)
	mock.lockAddBook.Unlock()
	mock.AddBookFunc(book)
}

// AddBookCalls gets all the calls that were made to AddBook.
// Check the length with:
//     len(mockedDataStore.AddBookCalls())
func (mock *DataStoreMock) AddBookCalls() []struct {
	Book *models.Book
} {
	var calls []struct {
		Book *models.Book
	}
	mock.lockAddBook.RLock()
	calls = mock.calls.AddBook
	mock.lockAddBook.RUnlock()
	return calls
}

// Close calls CloseFunc.
func (mock *DataStoreMock) Close(ctx context.Context) error {
	if mock.CloseFunc == nil {
		panic("DataStoreMock.CloseFunc: method is nil but DataStore.Close was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockClose.Lock()
	mock.calls.Close = append(mock.calls.Close, callInfo)
	mock.lockClose.Unlock()
	return mock.CloseFunc(ctx)
}

// CloseCalls gets all the calls that were made to Close.
// Check the length with:
//     len(mockedDataStore.CloseCalls())
func (mock *DataStoreMock) CloseCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockClose.RLock()
	calls = mock.calls.Close
	mock.lockClose.RUnlock()
	return calls
}

// Init calls InitFunc.
func (mock *DataStoreMock) Init(in1 config.MongoConfig) error {
	if mock.InitFunc == nil {
		panic("DataStoreMock.InitFunc: method is nil but DataStore.Init was just called")
	}
	callInfo := struct {
		In1 config.MongoConfig
	}{
		In1: in1,
	}
	mock.lockInit.Lock()
	mock.calls.Init = append(mock.calls.Init, callInfo)
	mock.lockInit.Unlock()
	return mock.InitFunc(in1)
}

// InitCalls gets all the calls that were made to Init.
// Check the length with:
//     len(mockedDataStore.InitCalls())
func (mock *DataStoreMock) InitCalls() []struct {
	In1 config.MongoConfig
} {
	var calls []struct {
		In1 config.MongoConfig
	}
	mock.lockInit.RLock()
	calls = mock.calls.Init
	mock.lockInit.RUnlock()
	return calls
}
