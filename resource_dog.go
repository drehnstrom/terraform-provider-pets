package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

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
func resourceDogRead(d *schema.ResourceData, m interface{}) error {
	return nil
}
func resourceDogUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

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
