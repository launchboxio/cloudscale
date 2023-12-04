import React, {useEffect, useState} from "react";
import axios from "axios";
import {Button, Card, FormGroup, HTMLTable, InputGroup, Intent, Popover} from "@blueprintjs/core";

const Listener = () => {
  const [listeners, setListeners] = useState([])
  const [newListener, setNewListener] = useState({})
  const [modalOpen, setModalOpen] = useState(false)

  useEffect(() => {
    axios.get("/listeners").then((res) => setListeners(res.data.listeners))
  }, [])

  const handleSubmit = (event) => {
    event.preventDefault()

    newListener.port = Number(newListener.port)
    axios.post("/listeners", newListener).then((res) => {
      setListeners(listeners.concat([res.data.listener]))
      setNewListener({})
      setModalOpen(false)
    })
  }

  const disable = (listenerId) => {
    axios.post(`/listeners/${listenerId}`, {
      enabled: false,
    })
  }

  return (
    <>
      <Card>
        <h4>Listeners</h4>
        <HTMLTable>
          <thead>
          <tr>
            <th>ID</th>
            <th>Name</th>
            <th>IP Address</th>
            <th>Port</th>
            <th></th>
          </tr>
          </thead>
          <tbody>
          {listeners.map((item) => {
            return (
              <tr>
                <td>{item.id}</td>
                <td>{item.name}</td>
                <td>{item.ip_address}</td>
                <td>{item.port}</td>
                <td><button>Delete</button></td>
              </tr>
            )
          })}
          </tbody>
        </HTMLTable>
      </Card>

      <Popover
        enforceFocus={false}
        isOpen={modalOpen}
        content={
          <div>
            <form onSubmit={handleSubmit}>
              <FormGroup
                helperText="The friendly name of your listener"
                label="Name"
                labelFor="name-input"
                labelInfo="(required)"
              >
                <InputGroup id="name-input" placeholder="My listener" onChange={(event) => {
                  setNewListener({
                    ...newListener,
                    name: event.target.value
                  })
                }}/>
              </FormGroup>
              <FormGroup
                helperText="The IP Address to bind to. Use 0.0.0.0 to bind to all networks"
                label="IP Address"
                labelFor="ip-address-input"
              >
                <InputGroup id="ip-address-input" placeholder="0.0.0.0" onChange={(event) => {
                  setNewListener({
                    ...newListener,
                    ip_address: event.target.value
                  })
                }}/>
              </FormGroup>
              <FormGroup
                helperText="The port to bind to"
                label="Port"
                labelFor="port-input"
                labelInfo="(required)"
              >
                <InputGroup id="port-input" onChange={(event) => {
                  setNewListener({
                    ...newListener,
                    port: event.target.value
                  })
                }}/>
              </FormGroup>
              <Button intent={Intent.PRIMARY} text={"Submit"} type={"submit"} />
              <Button intent={Intent.NONE} text={"Cancel"} onClick={() => {
                setNewListener({})
                setModalOpen(false)
              }} />
            </form>
          </div>
        }
      >

        <Button intent={Intent.PRIMARY} text={"Create New Listener"} tabIndex={0} onClick={() => setModalOpen(true)} />
      </Popover>
    </>
  )
}

export default Listener
