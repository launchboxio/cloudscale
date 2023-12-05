import * as React from "react"
import {Form} from "react-router-dom";
import {FormGroup, HTMLSelect, InputGroup, Switch} from "@blueprintjs/core";

export default () => {
  return (
    <Form method="post">
      <FormGroup
        label={"Type"}
        labelFor={"type-input"}
      >
        <HTMLSelect
          options={[{
            label: "HTTP (Layer 7)", value: "layer7",
          }, {
            label: "TCP (Layer 4)", value: "layer4"
          }]}
          value={"layer4"}
        />
      </FormGroup>
      <FormGroup
        helperText="The friendly name of your listener"
        label="Name"
        labelFor="name-input"
        labelInfo="(required)"
      >
        <InputGroup id="name-input" placeholder="My listener" name={"name"}/>
      </FormGroup>
      <FormGroup
        helperText="The IP Address to bind to. Use 0.0.0.0 to bind to all networks"
        label="IP Address"
        labelFor="ip-address-input"
        labelInfo={"(required)"}
      >
        <InputGroup id="ip-address-input" placeholder="0.0.0.0" name={"ip_address"}/>
      </FormGroup>
      <FormGroup
        helperText="The port to bind to"
        label="Port"
        labelFor="port-input"
        labelInfo="(required)"
      >
        <InputGroup id="port-input" name={"port"}/>
      </FormGroup>
      <button type="submit">New</button>
    </Form>
  )
}
