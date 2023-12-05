import * as React from "react"
import {Form} from "react-router-dom";
import {FormGroup, HTMLSelect, InputGroup, Switch} from "@blueprintjs/core";

export default () => {
  return (
    <Form method="post">
      <FormGroup
        helperText="The name of the target group"
        label="Name"
        labelFor="name-input"
        labelInfo="(required)"
      >
        <InputGroup id="name-input" placeholder="Target Group name" name={"name"}/>
      </FormGroup>
      <button type="submit">New</button>
    </Form>
  )
}
