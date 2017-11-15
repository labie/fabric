/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package binary

import (
	"archive/tar"
	"io/ioutil"
	"strings"
	cutil "github.com/hyperledger/fabric/core/container/util"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// Platform for the CAR type
type Platform struct {
}

// ValidateSpec validates the chaincode specification for CAR types to satisfy
// the platform interface.  This chaincode type currently doesn't
// require anything specific so we just implicitly approve any spec
func (binPlatform *Platform) ValidateSpec(spec *pb.ChaincodeSpec) error {
	return nil
}

func (binPlatform *Platform) ValidateDeploymentSpec(cds *pb.ChaincodeDeploymentSpec) error {
	// CAR platform will validate the code package within chaintool
	return nil
}

func (binPlatform *Platform) GetDeploymentPayload(spec *pb.ChaincodeSpec) ([]byte, error) {
	// path could be relative path or absolute path
	return ioutil.ReadFile(spec.ChaincodeId.Path)
}

func (binPlatform *Platform) GenerateDockerfile(cds *pb.ChaincodeDeploymentSpec) (string, error) {

	var buf []string

	// we assume that the chain code is written in golang
	buf = append(buf, "FROM "+cutil.GetDockerfileFromConfig("chaincode.golang.runtime"))
	buf = append(buf, "ADD binpackage.tar /usr/local/bin")

	dockerFileContents := strings.Join(buf, "\n")

	return dockerFileContents, nil
}

func (binPlatform *Platform) GenerateDockerBuild(cds *pb.ChaincodeDeploymentSpec, tw *tar.Writer) error {

	// simply tar provided chaincode binary package
	return cutil.WriteBytesToPackage("binpackage.tar", cds.CodePackage, tw)
}
