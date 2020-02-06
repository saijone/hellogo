package main
import ( 
    "fmt"
	"github.com/urfave/cli"
    "os"
    "os/exec"
)

// sudo iscsiadm -m node -T iqn.2015-12.com.oracleiaas:3a8537c2-c112-4bec-ae2d-0c2672387fe7 -p 169.254.2.5:3260 -u
// sudo iscsiadm -m node -o delete -T iqn.2015-12.com.oracleiaas:3a8537c2-c112-4bec-ae2d-0c2672387fe7 -p 169.254.2.5:3260
func detach(c *cli.Context) error {
    iqn := c.Args().Get(0);
    prt := c.Args().Get(1);
    fmt.Printf("   iqn %q", iqn)
    fmt.Printf("portal %q", prt)
    printOut(exec.Command("iscsiadm", "-m", "node", "-T", iqn, "-p", prt, "-u"))
    printOut(exec.Command("iscsiadm", "-m", "node", "-o", "delete", "-T", iqn, "-p", prt))
    return nil
}

func printOut(cmd *exec.Cmd) {
  out, err := cmd.CombinedOutput()
  if err != nil {
    fmt.Printf("%v", err)
  }
  if out != nil {
    fmt.Printf("%s", out)
  }
}
  
// sudo iscsiadm -m node -o new -T iqn.2015-12.com.oracleiaas:3a8537c2-c112-4bec-ae2d-0c2672387fe7 -p 169.254.2.3:3260
// sudo iscsiadm -m node -o update -T iqn.2015-12.com.oracleiaas:3a8537c2-c112-4bec-ae2d-0c2672387fe7 -n node.startup -v automatic
// sudo iscsiadm -m node -T iqn.2015-12.com.oracleiaas:3a8537c2-c112-4bec-ae2d-0c2672387fe7 -p 169.254.2.3:3260 -l
  func attach(c *cli.Context) error {
      iqn := c.Args().Get(0);
      prt := c.Args().Get(1);
      fmt.Printf("   iqn %q", iqn)
      fmt.Printf("portal %q", prt)
      
      printOut(exec.Command("iscsiadm", "-m", "node", "-o", "new", "-T", iqn, "-p", prt))
      printOut(exec.Command("iscsiadm", "-m", "node", "-o", "update", "-T", iqn, "-n", "node.startup", "-v", "automatic"))
      printOut(exec.Command("iscsiadm", "-m", "node", "-T", iqn, "-p", prt, "-l"))
      return nil
  }
  
 var commands = []cli.Command{
  	{
  		Name:   "a",
  		Usage:  "a",
  		Action: attach,
  	},
  	{
  		Name:   "d",
  		Usage:  "d",
  		Action: detach,
  	},
  }
 
func main() {
	app := cli.NewApp()
	app.Name = "vmst"
	app.Usage = "VMS tools"
    app.Commands = commands
    err := app.Run(os.Args)
    if err != nil {
    //   log.Fatal(err)
      fmt.Printf("%v", err)
    }
}