// MIT License
//
// Copyright (c) 2024 - Yuuki Takezawa
// Copyright (c) 2022 - present Open Swoole Group
// Copyright (c) 2018 SpiralScout
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package language

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

type PHPCode struct {
	PHP
}

func NewPHPCode() *PHPCode {
	return &PHPCode{PHP{}}
}

func (p *PHPCode) Generate(request *pluginpb.CodeGeneratorRequest) *pluginpb.CodeGeneratorResponse {
	response := &pluginpb.CodeGeneratorResponse{}
	for _, file := range request.ProtoFile {
		for _, service := range file.Service {
			response.File = append(response.File, p.generateInterface(request, file, service))
			response.File = append(response.File, p.generateService(request, file, service))
			response.File = append(response.File, p.generateClient(request, file, service))
		}
	}
	return response
}

func (p *PHPCode) generateClient(
	req *pluginpb.CodeGeneratorRequest,
	file *descriptorpb.FileDescriptorProto,
	service *descriptorpb.ServiceDescriptorProto) *pluginpb.CodeGeneratorResponse_File {

	nic := NewClientCode(req, file, service, NewNamespace(p.PHP, req, file, service))
	code, _ := nic.Body()
	i := &Client{p: p.PHP}
	return &pluginpb.CodeGeneratorResponse_File{
		Name:    proto.String(i.Filename(file, service.Name)),
		Content: proto.String(code),
	}
}

func (p *PHPCode) generateInterface(
	req *pluginpb.CodeGeneratorRequest,
	file *descriptorpb.FileDescriptorProto,
	service *descriptorpb.ServiceDescriptorProto) *pluginpb.CodeGeneratorResponse_File {

	nic := NewInterfaceCode(req, file, service, NewNamespace(p.PHP, req, file, service))
	code, _ := nic.Body()
	i := &InterfaceName{p: p.PHP}
	return &pluginpb.CodeGeneratorResponse_File{
		Name:    proto.String(i.Filename(file, service.Name)),
		Content: proto.String(code),
	}
}

func (p *PHPCode) generateService(
	req *pluginpb.CodeGeneratorRequest,
	file *descriptorpb.FileDescriptorProto,
	service *descriptorpb.ServiceDescriptorProto) *pluginpb.CodeGeneratorResponse_File {

	nic := NewServiceCode(req, file, service, NewNamespace(p.PHP, req, file, service))
	code, _ := nic.Body()
	i := &ServiceName{p: p.PHP}
	return &pluginpb.CodeGeneratorResponse_File{
		Name:    proto.String(i.Filename(file, service.Name)),
		Content: proto.String(code),
	}
}
