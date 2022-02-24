package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func CrudFirebase() {
	// Firestore初期化
	ctx := context.Background()
	sa := option.WithCredentialsFile("/Users/msonobe/DevSpace/Go/simple_web/path/to/serveiceAccount.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// CREATE:
	// データ追加
	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
		"first":  "Ada",
		"middle": "Mathison",
		"last":   "Lovelace",
		"born":   1815,
	})
	if err != nil {
		log.Fatalf("Failsed adding alovelace: %v", err)
	}

	// ドキュメント追加
	_, err = client.Collection("users").Doc("user1").Set(ctx, map[string]interface{}{
		"first":  "Ada",
		"middle": "Mathison",
		"last":   "Lovelace",
		"born":   1815,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

	// READ: データ読み取り
	iter := client.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}

	// UPDATE:
	// フィールド更新
	_, updateError := client.Collection("users").Doc("user1").Set(ctx, map[string]interface{}{
		"first": "Yeah",
	}, firestore.MergeAll)
	if updateError != nil {
		log.Printf("An error has occured: %s", err)
	}

	// フィールド削除
	_, fDeleteError := client.Collection("users").Doc("user1").Update(ctx, []firestore.Update{
		{
			Path:  "middle",
			Value: firestore.Delete,
		},
	})
	if fDeleteError != nil {
		log.Printf("An error has occured: %s", err)
	}

	// DELETE:
	// ドキュメント削除
	_, dDeleteError := client.Collection("users").Doc("user1").Delete(ctx)
	if dDeleteError != nil {
		log.Printf("An error has occured: %s", err)
	}

	// コレクション内のドキュメントを削除
	iter = client.Collection("user").Limit(10).Documents(ctx)
	numDeleted := 0
	batch := client.Batch()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("An error has occured: %s", err)
		}

		batch.Delete(doc.Ref)
		numDeleted++
	}

	// 切断
	defer client.Close()
}
