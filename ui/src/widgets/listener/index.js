import React, {useEffect, useState} from "react";
import axios from "axios";
import {Card} from "@blueprintjs/core";

const Listener = () => {
  const [listeners, setListeners] = useState([])
  const [newListener, setNewListener] = useState({})

  useEffect(() => {
    axios.get("/listeners").then((res) => setListeners(res.data.listeners))
  }, [])

  const handleSubmit = (event) => {
    event.preventDefault()

    newListener.port = Number(newListener.port)
    axios.post("/listeners", newListener).then((res) => {
      setListeners(listeners.concat([res.data.listener]))
      setNewListener({})
    })
  }
  return (
    <Card>

    </Card>
    <>
      <h5>Listeners</h5>
      <table>
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
      </table>

      <h5>Create new listener</h5>
      <form onSubmit={handleSubmit}>
        <input type="text" onChange={(event) => {
          setNewListener({
            ...newListener,
            name: event.target.value
          })
        }} value={newListener.name} />
        <input type="text" onChange={(event) => {
          setNewListener({
            ...newListener,
            ip_address: event.target.value
          })
        }} value={newListener.ip_address} />
        <input type="text" onChange={(event) => {
          setNewListener({
            ...newListener,
            port: event.target.value
          })
        }} value={newListener.port} />
        <input  type="submit"/>
      </form>
    </>
  )
}

export default Listener
