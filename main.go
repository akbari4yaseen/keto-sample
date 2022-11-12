package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	acl "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

func checkAcess() {
	conn, err := grpc.Dial("88.208.199.201:4466", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}

	client := acl.NewCheckServiceClient(conn)

	res, err := client.Check(context.Background(), &acl.CheckRequest{
		Namespace: "shifts-namespace",
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

func main() {
	conn, err := grpc.Dial("88.208.199.201:4467", grpc.WithInsecure())
	if err != nil {
		panic("Encountered error: " + err.Error())
	}

	client := acl.NewWriteServiceClient(conn)

	_, err = client.TransactRelationTuples(context.Background(), &acl.TransactRelationTuplesRequest{
		RelationTupleDeltas: []*acl.RelationTupleDelta{
			{
				Action: acl.RelationTupleDelta_ACTION_INSERT,
				RelationTuple: &acl.RelationTuple{
					Namespace: "shifts-namespace",
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
	checkAcess()
}
