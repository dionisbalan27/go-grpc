package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "latihan2.grpc/peserta"
)

type pesertaService struct {
	pb.UnimplementedPesertaServer
	pesertas []*pb.PesertaData
}

func (p *pesertaService) FindPesertaByNis(ctx context.Context, in *pb.PesertaRequest) (*pb.PesertaData, error) {
	for _, peserta := range p.pesertas {
		if peserta.Nis == in.Nis {
			return peserta, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "--Peserta tidak ditemukan--")
}

func newPesertaService() *pesertaService {
	var pesertas []*pb.PesertaData
	pesertas = append(pesertas, &pb.PesertaData{Nis: "123", Name: "Toni", Address: "Jakarta"})
	pesertas = append(pesertas, &pb.PesertaData{Nis: "124", Name: "Yuni", Address: "Bandung"})

	ps := &pesertaService{}
	ps.pesertas = pesertas
	return ps
}

func main() {
	lis, err := net.Listen("tcp", ":1200")
	if err != nil {
		log.Fatalf("error when listening : %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPesertaServer(s, newPesertaService())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("error when serve : %v", err)
	}
}
