package db

import (
	"context"
	"dicoding/Bookskette/backend/utils/random"
	"testing"

	"github.com/stretchr/testify/require"
)

func createInsertBook(t *testing.T) Book {
	var arg InsertBookParams = InsertBookParams{
		Title:  random.RandomTitle(),
		Author: random.RandomAuthor(),
		IsRead: random.RandomIsRead(),
		Year:   random.RandomYear(),
	}

	book, err := testQueries.InsertBook(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, book)

	require.Equal(t, arg.Title, book.Title)
	require.Equal(t, arg.Author, book.Author)
	require.Equal(t, arg.IsRead, book.IsRead)
	require.Equal(t, arg.Year, book.Year)

	require.NotZero(t, book.ID)
	require.NotZero(t, book.CreatedAt)

	return book
}

func TestInsertBook(t *testing.T) {
	createInsertBook(t)
}

func TestGetBook(t *testing.T) {
	createdBook := createInsertBook(t)

	book, err := testQueries.GetBook(context.Background(), createdBook.ID)
	require.NoError(t, err)
	require.NotEmpty(t, book)

	require.Equal(t, createdBook.Title, book.Title)
	require.Equal(t, createdBook.Author, book.Author)
	require.Equal(t, createdBook.IsRead, book.IsRead)
	require.Equal(t, createdBook.Year, book.Year)
}

func TestListBook(t *testing.T) {
	ammount := 3
	for i := 0; i < ammount; i++ {
		createInsertBook(t)
	}

	listBook, err := testQueries.ListBook(context.Background())
	require.NoError(t, err)

	for i := 0; i < len(listBook); i++ {
		require.NotEmpty(t, listBook[i])
	}

}

func TestUpdateBook(t *testing.T) {
	createdBook := createInsertBook(t)

	arg := UpdateBookParams{
		ID:     createdBook.ID,
		Title:  random.RandomTitle(),
		Author: random.RandomAuthor(),
		IsRead: random.RandomIsRead(),
		Year:   random.RandomYear(),
	}

	updatedBook, err := testQueries.UpdateBook(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedBook)

	require.Equal(t, arg.Title, updatedBook.Title)
	require.Equal(t, arg.Author, updatedBook.Author)
	require.Equal(t, arg.IsRead, updatedBook.IsRead)
	require.Equal(t, arg.Year, updatedBook.Year)
}

func TestToggleBookStatus(t *testing.T) {
	createdBook := createInsertBook(t)

	updatedBook, err := testQueries.ToggleBookStatus(context.Background(), createdBook.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedBook)

	require.Equal(t, !createdBook.IsRead, updatedBook.IsRead)

}

func TestDeleteBook(t *testing.T) {
	createdBook := createInsertBook(t)

	deletedBook, err := testQueries.DeleteBook(context.Background(), createdBook.ID)
	require.NoError(t, err)
	require.NotEmpty(t, deletedBook)

	require.Equal(t, createdBook.ID, deletedBook.ID)

	deletedBook, err = testQueries.GetBook(context.Background(), createdBook.ID)
	require.Empty(t,deletedBook)
}

