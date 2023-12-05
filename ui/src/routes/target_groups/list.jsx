import * as React from "react"
import {Button, Card, HTMLTable, Intent} from "@blueprintjs/core";
import {Link, useLoaderData} from "react-router-dom";

const List = () => {
  const { target_groups } = useLoaderData();
  return (
    <>
      <Card>
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
          {target_groups.map((item) => {
            return (
              <tr>
                <td>{item.id}</td>
                <td>{item.name}</td>
                <td>{item.attachments.length}</td>
                <td>
                  <Link to={`/target_groups/${item.id}`}>
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
      <Link to={"/target_groups/new"}>Create new TargetGroup</Link>
    </>
  )
}

export default List
