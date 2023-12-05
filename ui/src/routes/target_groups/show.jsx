import * as React from "react";
import {Form, useLoaderData} from "react-router-dom";
import {Button, Card, FormGroup, InputGroup, Intent} from "@blueprintjs/core";

export default () => {
  const { target_group } = useLoaderData()

  return (
    <>
      <h1>{target_group.name}</h1>
      <Card>
        <h4>Attachments</h4>
        <ul>
          {target_group.attachments.map((item) => {
            return (
              <li>
                {item.ip_address}:{item.port}
              </li>
            )
          })}
        </ul>
      </Card>
      <Card>
        <h4>Add Upstream</h4>
        <Form method={"post"}>
          <FormGroup
            label="IP Address"
            labelFor={`ip_address`}
          >
            <InputGroup id={"ip_address"} name={"ip_address"}/>
          </FormGroup>
          <FormGroup
            label="Port"
            labelFor={"port"}
          >
            <InputGroup id={"port"} name={"port"} />
            <Button intent={Intent.NONE} text={"Add Attachment"} type={"submit"} />
          </FormGroup>
        </Form>
      </Card>
    </>
  )
}
