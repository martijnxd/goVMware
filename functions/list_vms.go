package functions

import (
	"context"
	"fmt"

	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
)

func ListVMS() {

	Run(func(ctx context.Context, c *vim25.Client) error {
		// Create view of VirtualMachine objects

		m := view.NewManager(c)

		v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{"VirtualMachine"}, true)
		if err != nil {
			return err
		}

		defer v.Destroy(ctx)

		// Retrieve summary property for all machines
		// Reference: http://pubs.vmware.com/vsphere-60/topic/com.vmware.wssdk.apiref.doc/vim.VirtualMachine.html
		var vms []mo.VirtualMachine
		err = v.Retrieve(ctx, []string{"VirtualMachine"}, []string{"summary"}, &vms)
		if err != nil {
			return err
		}

		// Print summary per vm (see also: govc/vm/info.go)

		for _, vm := range vms {
			fmt.Printf("%s: %s: %s\n", vm.Summary.Config.Name, vm.Summary.Config.GuestFullName, vm.Summary.Guest.HostName)
		}

		return nil
	})
}
