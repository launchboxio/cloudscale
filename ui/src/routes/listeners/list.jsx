import * as React from "react"
import {Card, HTMLTable} from "@blueprintjs/core";
import {Link, useLoaderData} from "react-router-dom";

const List = () => {
  const { listeners } = useLoaderData();
  console.log(listeners)
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
                <td>
                  <Link to={`/listeners/${item.id}`}>
                    View
                  </Link>
                  <button>Delete</button>
                </td>
              </tr>
            )
          })}
          </tbody>
        </HTMLTable>
      </Card>
      <Link to={"/listeners/new"}>Create new Listener</Link>
    </>
  )
}

export default List
