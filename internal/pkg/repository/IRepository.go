package repository

type IRepository interface {
	Add() error
	Get() (string, error)
	CheckIfFullUrlExists() (bool, error)
	CheckIfTinyUrlExists() (bool, error)
}
