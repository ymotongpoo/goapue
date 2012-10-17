/**
 * p.14
 */

package main

import (
	"os/user"
	"fmt"
)


func main() {
	u, err := user.Current()
	if err != nil {
		fmt.Errorf("error")
	}
	fmt.Printf("uid = %s, gid = %s\n", u.Uid, u.Gid)
}