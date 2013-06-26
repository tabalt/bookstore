package bookstore

import (
    "fmt"
    "github.com/tabalt/appgo"
)

func NewBook() *Book {
    return &Book{}
}
type BookInfo struct {
    Id int
    Name string
    Price float64
}

type Book struct {
    appgo.App
    bookshelf []BookInfo
}

func (this *Book) Add() {
    fmt.Println("Add Book")
}

func (this *Book) List() {
    fmt.Println("List Book")
}

func (this *Book) Find() {
    fmt.Println("Find Book")
}