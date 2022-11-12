package main

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/grpc"

	acl "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

func CheckAccess() {
	conn, err := grpc.Dial(os.Getenv("KETO_READ_API"), grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}

	client := acl.NewCheckServiceClient(conn)

	res, err := client.Check(context.Background(), &acl.CheckRequest{
		Namespace: os.Getenv("KETO_NAMESPACE"),
		Object:    "02y_15_4w350m3",
		Relation:  "decypher",
		Subject:   acl.NewSubjectID("yaseen"),
	})
	if err != nil {
		panic(err.Error())
	}

	if res.Allowed {
		fmt.Println("Allowed")
		return
	}
	fmt.Println("Denied")
}

func GrantAccess() {
	conn, err := grpc.Dial(os.Getenv("KETO_WRITE_API"), grpc.WithInsecure())
	if err != nil {
		panic("Encountered error: " + err.Error())
	}

	client := acl.NewWriteServiceClient(conn)

	_, err = client.TransactRelationTuples(context.Background(), &acl.TransactRelationTuplesRequest{
		RelationTupleDeltas: []*acl.RelationTupleDelta{
			{
				Action: acl.RelationTupleDelta_ACTION_INSERT,
				RelationTuple: &acl.RelationTuple{
					Namespace: os.Getenv("KETO_NAMESPACE"),
					Object:    "02y_15_4w350m3",
					Relation:  "decypher",
					Subject:   acl.NewSubjectID("john"),
				},
			},
		},
	})

	if err != nil {
		panic("Encountered error: " + err.Error())
	}

	fmt.Println("Successfully created tuple")
}

func main() {
	GrantAccess()
	CheckAccess()
}
