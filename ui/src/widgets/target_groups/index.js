import React, {useEffect, useState} from "react";
import axios from "axios";
import {Button, FormGroup, HTMLTable, InputGroup, Intent, Popover, TextArea} from "@blueprintjs/core";

const TargetGroups = () => {
  const [targetGroups, setTargetGroups] = useState([])
  const [newTargetGroup, setNewTargetGroup] = useState({
    attachments: []
  })
  const [currentState, setCurrentState] = useState('list')
  useEffect(() => {
    axios.get("/target_groups").then((res) => setTargetGroups(res.data.target_groups || []))
  }, [])

  const addAttachment = () => {
    setNewTargetGroup({
      ...newTargetGroup,
      attachments: newTargetGroup.attachments.concat([{}])
    })
  }

  const handleCreate = (event) => {
    event.preventDefault()
    newTargetGroup.attachments.map((item) => {
      item.port = Number(item.port)
      return item
    })
    axios.post("/target_groups", newTargetGroup).then((res) => {
      setNewTargetGroup({})
      setCurrentState("list")
      setTargetGroups(targetGroups.concat([res.data.target_group]))
    })
  }

  const handleAttachmentChange = (idx, property, value) => {
    let { attachments } = newTargetGroup
    attachments[idx][property] = value
    setNewTargetGroup({
      ...newTargetGroup,
      attachments
    })
  }

  return (
    <>
      {currentState === 'list' && (
        <>
          <h4>Target Groups</h4>
          <HTMLTable>
            <thead>
            <tr>
              <th>ID</th>
              <th>Name</th>
              <th>Attachments</th>
              <th></th>
            </tr>
            </thead>
            <tbody>
            {targetGroups.map((item) => {
              return (
                <tr>
                  <td>{item.id}</td>
                  <td>{item.name}</td>
                  <td>{item.attachments.length}</td>
                </tr>
              )
            })}
            </tbody>
          </HTMLTable>
          <Button intent={Intent.PRIMARY} text={"Create New Target Group"} tabIndex={0} onClick={() => setCurrentState('creating')} />
        </>
    )}

      {currentState === 'creating' && (
        <div>
          <form onSubmit={handleCreate}>
            <FormGroup
              helperText="The name of the target group"
              label="Name"
              labelFor="name-input"
              labelInfo="(required)"
            >
              <InputGroup id="name-input" placeholder="Target Group name" onChange={(event) => {
                setNewTargetGroup({
                  ...newTargetGroup,
                  name: event.target.value
                })
              }}/>
            </FormGroup>
            {newTargetGroup.attachments?.map((item, index) => {
              return (
                <>
                  <FormGroup
                    label="IP Address"
                    labelFor={`att-${index}-ipaddr`}
                  >
                    <InputGroup id={`att-${index}-ipaddr`} onChange={(event) => {
                      handleAttachmentChange(index, "ip_address", event.target.value)
                    }}/>
                  </FormGroup>
                  <FormGroup
                    label="Port"
                    labelFor={`att-${index}-port`}
                  >
                    <InputGroup id={`att-${index}-port`} onChange={(event) => {
                      handleAttachmentChange(index, "port", event.target.value)
                    }}/>
                  </FormGroup>
                </>
              )
            })}
            <Button intent={Intent.NONE} text={"Add Attachment"} onClick={() => {
              addAttachment()
            }} />
            <Button intent={Intent.PRIMARY} text={"Submit"} type={"submit"} />
            <Button intent={Intent.NONE} text={"Cancel"} onClick={() => {
              setNewTargetGroup({})
              setCurrentState("list")
            }} />
          </form>
        </div>
      )}
    </>
  )
}

export default TargetGroups
