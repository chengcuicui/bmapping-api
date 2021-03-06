package models

import (
	"fmt"
	"testing"

	"github.com/relax-space/go-kit/test"
)

func Test_GreenStore_GetEIdByCode(t *testing.T) {
	has, eId, store, err := GetEIdByCode(ctx, "AA01", 1)
	fmt.Println(has, eId, *store)
	test.Ok(t, err)
}

func Test_GreenStore_InsertMany(t *testing.T) {
	d := GreenStore{
		Code: "0002",
		Name: "绿色商店3",
	}
	var storegroups []GreenStore
	storegroups = append(storegroups, d)

	err := (GreenStore{}).InsertMany(ctx, &storegroups)
	test.Ok(t, err)
}

func Test_GreenStore_GetById(t *testing.T) {
	has, greenStore, err := GreenStore{}.GetById(ctx, 2)
	fmt.Println(has, greenStore)
	test.Ok(t, err)
}

func Test_GreenStore_GetAll(t *testing.T) {
	count, items, err := GreenStore{}.GetAll(ctx, nil, nil, 0, 3)
	fmt.Println(count, items)
	test.Ok(t, err)
}

func Test_GreenStore_Update(t *testing.T) {
	d := &GreenStore{
		Name: "流通百货",
	}
	err := d.Update(ctx, 2)
	test.Ok(t, err)
}

func Test_GreenStore_InsertOne(t *testing.T) {
	d := &GreenStore{
		Code: "CR03",
		Name: "流通百货3",
	}
	err := d.InsertOne(ctx)
	test.Ok(t, err)
}

func Test_GreenStore_Delete(t *testing.T) {
	err := GreenStore{}.Delete(ctx, 1)
	test.Ok(t, err)
}


//greenStoreGroup
func Test_GreenStoreGroup_InsertMany(t *testing.T) {
	d := GreenStoreGroup{
		Code: "CR01",
		Name: "衣服裤子",
	}
	var storegroups []GreenStoreGroup
	storegroups = append(storegroups, d)

	err := (GreenStoreGroup{}).InsertMany(ctx, &storegroups)
	test.Ok(t, err)
}

func Test_GreenStoreGroup_GetById(t *testing.T) {
	has, greenStoreGroup, err := GreenStoreGroup{}.GetById(ctx, 1)
	fmt.Println(has, greenStoreGroup)
	test.Ok(t, err)
}

func Test_GreenStoreGroup_GetAll(t *testing.T) {
	count, items, err := GreenStoreGroup{}.GetAll(ctx, nil, nil, 0, 2)
	fmt.Println(count, items)
	test.Ok(t, err)
}

func Test_GreenStoreGroup_Update(t *testing.T) {
	d := &GreenStoreGroup{
		Name: "流通百货",
	}
	err := d.Update(ctx, 2)
	test.Ok(t, err)
}

func Test_GreenStoreGroup_InsertOne(t *testing.T) {
	d := &GreenStoreGroup{
		Code: "CR03",
		Name: "流通百货3",
	}
	err := d.InsertOne(ctx)
	test.Ok(t, err)
}

func Test_GreenStoreGroup_Delete(t *testing.T) {
	err := GreenStoreGroup{}.Delete(ctx, 1)
	test.Ok(t, err)
}
