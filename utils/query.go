package utils

import (
	"context"
	"database/sql"
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

	var id uint64
	err = stmt.QueryRowContext(ctx, requests...).Scan(&id)
	if err != nil {
		// Handle specific errors (e.g., pq.Error)
		return nil, err
	}
	return &id, nil
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
	log.Println(strings.ReplaceAll(query.ToString(), "$1", fmt.Sprintf("%d", id)))

	err := QueryValidation(query, selectQuery)
	if err != nil {
		return err
	}
	stmt, err := db.PrepareContext(ctx, query.ToString())
	if err != nil {
		return err
	}
	defer stmt.Close()

	var ptrs []interface{}

	typ := reflect.TypeOf(result).Elem()

	for i := 0; i < typ.NumField(); i++ {
		field := reflect.ValueOf(result).Elem().Field(i)
		ptrs = append(ptrs, field.Addr().Interface())
	}

	err = stmt.QueryRowContext(ctx, id).Scan(ptrs...)
	if err != nil {
		return err
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
