package main

import (
	"context"
	"testing"

	pb "Nftsale_cash_delivery/src/productcatalogservice/genproto"
	"github.com/golang/protobuf/proto"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestServer(t *testing.T) {
	ctx := context.Background()
	addr := run(port)
	conn, err := grpc.DialContext(ctx, addr,
		grpc.WithInsecure())

	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewProductCatalogServiceClient(conn)
	res, err := client.ListProducts(ctx, &pb.Empty{})
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(res.Products, parseCatalog(), cmp.Comparer(proto.Equal)); diff != "" {
		t.Error(diff)
	}

	got, err := client.GetProduct(ctx, &pb.GetProductRequest{Id: "OLJCESPC7Z"})
	if err != nil {
		t.Fatal(err)
	}
	if want := parseCatalog()[0]; !proto.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
	_, err = client.GetProduct(ctx, &pb.GetProductRequest{Id: "N/A"})
	if got, want := status.Code(err), codes.NotFound; got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	sres, err := client.SearchProducts(ctx, &pb.SearchProductsRequest{Query: "sunglasses"})
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(sres.Results, []*pb.Product{parseCatalog()[0]}, cmp.Comparer(proto.Equal)); diff != "" {
		t.Error(diff)
	}
}
