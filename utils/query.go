package utils

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"log"
	"reflect"
	"strings"
)

const (
	selectQuery = `SELECT`
	insertQuery = `INSERT`
	updateQuery = `UPDATE`
	deleteQuery = `DELETE`
)

func QueryValidation(query queries.GoQuery, command string) error {
	queryStr := query.ToString()
	if !strings.Contains(queryStr, command) {
		return fmt.Errorf(`the query command is not valid`)
	}
	return nil
}

func Save[T any](ctx context.Context, db repositories_interfaces.Trigger, query queries.GoQuery, requests ...any) (*uint64, error) {
	err := QueryValidation(query, insertQuery)
	if err != nil {
		return nil, err
	}

	stmt, err := db.PrepareContext(ctx, query.ToString())
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute the insert statement
	result, err := stmt.ExecContext(ctx, requests...)
	if err != nil {
		return nil, err
	}

	// Retrieve the last inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	uid := uint64(id)
	return &uid, nil
}

func FindAll[T any](ctx context.Context, db *sql.DB, query queries.GoQuery) (result []*T, err error) {
	requests := new(T)
	err = QueryValidation(query, selectQuery)
	if err != nil {
		return nil, err
	}
	rows, err := db.QueryContext(ctx, query.ToString())
	if err != nil {
		return nil, err
	}

	typ := reflect.TypeOf(requests).Elem()
	var ptrs = make([]interface{}, typ.NumField())
	for i := range ptrs {
		ptrs[i] = reflect.New(typ.Field(i).Type).Interface()
	}

	for rows.Next() {
		request := reflect.New(typ).Elem()
		err := rows.Scan(ptrs...)
		if err != nil {
			return nil, err
		}
		for i, ptr := range ptrs {
			field := request.Field(i)
			field.Set(reflect.ValueOf(ptr).Elem())
		}
		result = append(result, request.Addr().Interface().(*T))
	}
	return result, nil
}

func Delete[T any](ctx context.Context, db repositories_interfaces.Trigger, query queries.GoQuery, id uint64) (err error) {
	err = QueryValidation(query, deleteQuery)
	if err != nil {
		return err
	}
	_, err = db.ExecContext(ctx, query.ToString(), id)
	if err != nil {
		return err
	}
	return nil
}

func FindByID[T any](ctx context.Context, db repositories_interfaces.Trigger, query queries.GoQuery, id uint64, result *T) error {
	// Log query for debugging
	log.Println("Executing query:", strings.ReplaceAll(query.ToString(), "?", fmt.Sprintf("%d", id)))

	// Validate the query to ensure it's a SELECT query
	err := QueryValidation(query, selectQuery)
	if err != nil {
		return fmt.Errorf("query validation failed: %w", err)
	}

	// Prepare the SQL statement
	stmt, err := db.PrepareContext(ctx, query.ToString())
	if err != nil {
		return fmt.Errorf("failed to prepare context: %w", err)
	}
	defer stmt.Close()

	// Use reflection to ensure the result is properly initialized
	if reflect.ValueOf(result).Kind() != reflect.Ptr || reflect.ValueOf(result).IsNil() {
		return fmt.Errorf("result argument must be a non-nil pointer")
	}

	// Get the underlying type and value of the result
	typ := reflect.TypeOf(result).Elem()
	val := reflect.ValueOf(result).Elem()

	var ptrs []interface{}

	// Create a slice of pointers to each field in the struct
	for i := 0; i < typ.NumField(); i++ {
		field := val.Field(i)
		ptrs = append(ptrs, field.Addr().Interface())
	}

	// Execute the query and scan the result into the struct fields
	err = stmt.QueryRowContext(ctx, id).Scan(ptrs...)
	if err != nil {
		// Handle no rows found
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("no rows found for ID %d: %w", id, err)
		}
		// Handle other scan errors
		return fmt.Errorf("failed to scan result for ID %d: %w", id, err)
	}

	return nil
}

func Update[T any](ctx context.Context, db *sql.DB, query queries.GoQuery, requests ...any) (err error) {
	err = QueryValidation(query, updateQuery)
	if err != nil {
		return err
	}
	_, err = db.ExecContext(ctx, query.ToString(), requests...)
	if err != nil {
		return err
	}
	return nil
}
