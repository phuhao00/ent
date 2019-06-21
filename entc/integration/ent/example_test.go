// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"log"
	"net/url"
	"os"
	"time"

	"fbc/ent/dialect"
	"fbc/lib/go/gremlin"
)

// endpoint for the database. In order to run the tests locally, run the following command:
//
//	 ENT_INTEGRATION_ENDPOINT="http://localhost:8182" go test -v
//
var endpoint *gremlin.Endpoint

func init() {
	if e, ok := os.LookupEnv("ENT_INTEGRATION_ENDPOINT"); ok {
		if u, err := url.Parse(e); err == nil {
			endpoint = &gremlin.Endpoint{u}
		}
	}
}

func ExampleCard() {
	if endpoint == nil {
		return
	}
	ctx := context.Background()
	conn, err := gremlin.NewClient(gremlin.Config{Endpoint: *endpoint})
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	client := NewClient(Driver(dialect.NewGremlin(conn)))

	// creating vertices for the card's edges.

	// create card vertex with its edges.
	c := client.Card.
		Create().
		SetNumber("string").
		SaveX(ctx)
	log.Println("card created:", c)

	// query edges.

	// Output:
}
func ExampleComment() {
	if endpoint == nil {
		return
	}
	ctx := context.Background()
	conn, err := gremlin.NewClient(gremlin.Config{Endpoint: *endpoint})
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	client := NewClient(Driver(dialect.NewGremlin(conn)))

	// creating vertices for the comment's edges.

	// create comment vertex with its edges.
	c := client.Comment.
		Create().
		SaveX(ctx)
	log.Println("comment created:", c)

	// query edges.

	// Output:
}
func ExampleFile() {
	if endpoint == nil {
		return
	}
	ctx := context.Background()
	conn, err := gremlin.NewClient(gremlin.Config{Endpoint: *endpoint})
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	client := NewClient(Driver(dialect.NewGremlin(conn)))

	// creating vertices for the file's edges.

	// create file vertex with its edges.
	f := client.File.
		Create().
		SetSize(1).
		SetName("string").
		SaveX(ctx)
	log.Println("file created:", f)

	// query edges.

	// Output:
}
func ExampleGroup() {
	if endpoint == nil {
		return
	}
	ctx := context.Background()
	conn, err := gremlin.NewClient(gremlin.Config{Endpoint: *endpoint})
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	client := NewClient(Driver(dialect.NewGremlin(conn)))

	// creating vertices for the group's edges.
	f0 := client.File.
		Create().
		SetSize(1).
		SetName("string").
		SaveX(ctx)
	log.Println("file created:", f0)
	u1 := client.User.
		Create().
		SetAge(1).
		SetName("string").
		SetLast("string").
		SetNickname("string").
		SetPhone("string").
		SaveX(ctx)
	log.Println("user created:", u1)
	gi3 := client.GroupInfo.
		Create().
		SetDesc("string").
		SetMaxUsers(1).
		SaveX(ctx)
	log.Println("groupinfo created:", gi3)

	// create group vertex with its edges.
	gr := client.Group.
		Create().
		SetActive(true).
		SetExpire(time.Now()).
		SetType("string").
		SetMaxUsers(1).
		SetName("string").
		AddFiles(f0).
		AddBlocked(u1).
		SetInfo(gi3).
		SaveX(ctx)
	log.Println("group created:", gr)

	// query edges.
	f0, err = gr.QueryFiles().First(ctx)
	if err != nil {
		log.Fatalf("failed querying files: %v", err)
	}
	log.Println("files found:", f0)

	u1, err = gr.QueryBlocked().First(ctx)
	if err != nil {
		log.Fatalf("failed querying blocked: %v", err)
	}
	log.Println("blocked found:", u1)

	gi3, err = gr.QueryInfo().First(ctx)
	if err != nil {
		log.Fatalf("failed querying info: %v", err)
	}
	log.Println("info found:", gi3)

	// Output:
}
func ExampleGroupInfo() {
	if endpoint == nil {
		return
	}
	ctx := context.Background()
	conn, err := gremlin.NewClient(gremlin.Config{Endpoint: *endpoint})
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	client := NewClient(Driver(dialect.NewGremlin(conn)))

	// creating vertices for the groupinfo's edges.

	// create groupinfo vertex with its edges.
	gi := client.GroupInfo.
		Create().
		SetDesc("string").
		SetMaxUsers(1).
		SaveX(ctx)
	log.Println("groupinfo created:", gi)

	// query edges.

	// Output:
}
func ExampleNode() {
	if endpoint == nil {
		return
	}
	ctx := context.Background()
	conn, err := gremlin.NewClient(gremlin.Config{Endpoint: *endpoint})
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	client := NewClient(Driver(dialect.NewGremlin(conn)))

	// creating vertices for the node's edges.
	n1 := client.Node.
		Create().
		SetValue(1).
		SaveX(ctx)
	log.Println("node created:", n1)

	// create node vertex with its edges.
	n := client.Node.
		Create().
		SetValue(1).
		SetNext(n1).
		SaveX(ctx)
	log.Println("node created:", n)

	// query edges.

	n1, err = n.QueryNext().First(ctx)
	if err != nil {
		log.Fatalf("failed querying next: %v", err)
	}
	log.Println("next found:", n1)

	// Output:
}
func ExamplePet() {
	if endpoint == nil {
		return
	}
	ctx := context.Background()
	conn, err := gremlin.NewClient(gremlin.Config{Endpoint: *endpoint})
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	client := NewClient(Driver(dialect.NewGremlin(conn)))

	// creating vertices for the pet's edges.

	// create pet vertex with its edges.
	pe := client.Pet.
		Create().
		SetName("string").
		SaveX(ctx)
	log.Println("pet created:", pe)

	// query edges.

	// Output:
}
func ExampleUser() {
	if endpoint == nil {
		return
	}
	ctx := context.Background()
	conn, err := gremlin.NewClient(gremlin.Config{Endpoint: *endpoint})
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	client := NewClient(Driver(dialect.NewGremlin(conn)))

	// creating vertices for the user's edges.
	c0 := client.Card.
		Create().
		SetNumber("string").
		SaveX(ctx)
	log.Println("card created:", c0)
	pe1 := client.Pet.
		Create().
		SetName("string").
		SaveX(ctx)
	log.Println("pet created:", pe1)
	f2 := client.File.
		Create().
		SetSize(1).
		SetName("string").
		SaveX(ctx)
	log.Println("file created:", f2)
	gr3 := client.Group.
		Create().
		SetActive(true).
		SetExpire(time.Now()).
		SetType("string").
		SetMaxUsers(1).
		SetName("string").
		SaveX(ctx)
	log.Println("group created:", gr3)
	u4 := client.User.
		Create().
		SetAge(1).
		SetName("string").
		SetLast("string").
		SetNickname("string").
		SetPhone("string").
		SaveX(ctx)
	log.Println("user created:", u4)
	u6 := client.User.
		Create().
		SetAge(1).
		SetName("string").
		SetLast("string").
		SetNickname("string").
		SetPhone("string").
		SaveX(ctx)
	log.Println("user created:", u6)
	pe7 := client.Pet.
		Create().
		SetName("string").
		SaveX(ctx)
	log.Println("pet created:", pe7)
	u8 := client.User.
		Create().
		SetAge(1).
		SetName("string").
		SetLast("string").
		SetNickname("string").
		SetPhone("string").
		SaveX(ctx)
	log.Println("user created:", u8)
	u10 := client.User.
		Create().
		SetAge(1).
		SetName("string").
		SetLast("string").
		SetNickname("string").
		SetPhone("string").
		SaveX(ctx)
	log.Println("user created:", u10)

	// create user vertex with its edges.
	u := client.User.
		Create().
		SetAge(1).
		SetName("string").
		SetLast("string").
		SetNickname("string").
		SetPhone("string").
		SetCard(c0).
		AddPets(pe1).
		AddFiles(f2).
		AddGroups(gr3).
		AddFriends(u4).
		AddFollowing(u6).
		SetTeam(pe7).
		SetSpouse(u8).
		SetParent(u10).
		SaveX(ctx)
	log.Println("user created:", u)

	// query edges.
	c0, err = u.QueryCard().First(ctx)
	if err != nil {
		log.Fatalf("failed querying card: %v", err)
	}
	log.Println("card found:", c0)

	pe1, err = u.QueryPets().First(ctx)
	if err != nil {
		log.Fatalf("failed querying pets: %v", err)
	}
	log.Println("pets found:", pe1)

	f2, err = u.QueryFiles().First(ctx)
	if err != nil {
		log.Fatalf("failed querying files: %v", err)
	}
	log.Println("files found:", f2)

	gr3, err = u.QueryGroups().First(ctx)
	if err != nil {
		log.Fatalf("failed querying groups: %v", err)
	}
	log.Println("groups found:", gr3)

	u4, err = u.QueryFriends().First(ctx)
	if err != nil {
		log.Fatalf("failed querying friends: %v", err)
	}
	log.Println("friends found:", u4)

	u6, err = u.QueryFollowing().First(ctx)
	if err != nil {
		log.Fatalf("failed querying following: %v", err)
	}
	log.Println("following found:", u6)

	pe7, err = u.QueryTeam().First(ctx)
	if err != nil {
		log.Fatalf("failed querying team: %v", err)
	}
	log.Println("team found:", pe7)

	u8, err = u.QuerySpouse().First(ctx)
	if err != nil {
		log.Fatalf("failed querying spouse: %v", err)
	}
	log.Println("spouse found:", u8)

	u10, err = u.QueryParent().First(ctx)
	if err != nil {
		log.Fatalf("failed querying parent: %v", err)
	}
	log.Println("parent found:", u10)

	// Output:
}