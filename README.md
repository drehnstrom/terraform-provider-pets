# This is a Demo Terraform Provider
The provider is named `pets` and creates a single resource called `dog`.

When creating a Terraform provider, the compiled program needs to be named:
```
terraform-provider-<PROVIDER NAME>
```
In this example, the name is defined in the `go.mod` file when defining the module.

```
module terraform-provider-pets
```

## Defining the Provider
The `main.go` file is just boilerplate code. Notice, the `main()` method returns the `Provider()` function which is defined in the `provider.go` file. 

```
package main
import (
        "github.com/hashicorp/terraform-plugin-sdk/plugin"
        "github.com/hashicorp/terraform-plugin-sdk/terraform"
)
func main() {
        plugin.Serve(&plugin.ServeOpts{
                ProviderFunc: func() terraform.ResourceProvider {
                        return Provider()
                },
        })
}
```

In the `provider.go` file, the resources are defined. In this case, there is just one resource, `dog`. Resources are defined in a map where the key is the name of the resource and the value is a function that defines the resource, in this case `resourceDog()`.

```
package main
import (
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
			ResourcesMap: map[string]*schema.Resource{
					"pets_dog": resourceDog(),
			},
	}
}
```

## Defining Resources
By convention, resources should be defined in a file named `resource_<resource-name>.go`. In this case, the file is `resource_dog.go`.

In that file, the `resourceDog()` function defines the CRUD operations for the resource, and the properties that are supported by the resource. 

```
func resourceDog() *schema.Resource {
	return &schema.Resource{
		Create: resourceDogCreate,
		Read:   resourceDogRead,
		Update: resourceDogUpdate,
		Delete: resourceDogDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"breed": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
```
The `Create()` function, creates the resource (obviously). In this example, it simply creates a file. It is important to note, that if no error is returned, Terraform asssumes everything worked. If an error is returned, it assumes it failed. 

Also, notice how the Id of the resource is set. This would be important in differentiating multiple instances of the same resource. 

```
func resourceDogCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	breed := d.Get("breed").(string)
	d.SetId(fmt.Sprintf("%v-the-%v", name, breed))

	filename:= fmt.Sprint(d.Id(), ".txt")
	message := []byte(fmt.Sprintf("Woof-I-am-%v-the-%v", name, breed))
	err := ioutil.WriteFile(filename, message, 0777)
	if err != nil {
		return err
	}

	return nil
}
```

The following function deletes the resource. As before, if everything works, return no error. 

```
func resourceDogDelete(d *schema.ResourceData, m interface{}) error {
	filename:= fmt.Sprint(d.Id(), ".txt")
	err := os.Remove(filename)
	if err != nil {
		return err
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return nil
}
```

To keep that example simple, I chose not to implement `Read()` or `Update()`. 

## Deploying your Provider

You need a folder to put your provider in. By convention this is in your home folder: `~/.terraform.d/plugins`. Make the folder like this (note the version and platform in the path): 

```
mkdir -p ~/.terraform.d/plugins/drehnstrom/providers/pets/0.0.1/darwin_amd64
```

Copy your provider into the folder. 

```
cp terraform-provider-pets ~/.terraform.d/plugins/drehnstrom/providers/pets/0.0.1/darwin_amd64
```
## Using your Provider

In a Terraform file, define your module. The source and version attributes need to match the path you deployed the module into above. 

```
terraform {
  required_providers {
    pets = {
      version = "~> 0.0.1"
      source  = "drehnstrom/providers/pets"
    }
  }
}
```

Now, you can create a resource defined by the provider. 

```
resource "pets_dog" "my-dog" {
  name     = "Noir"
  breed    = "Schnoodle"
}
```
