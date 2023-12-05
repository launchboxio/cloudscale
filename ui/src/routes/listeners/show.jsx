import * as React from "react";
import {useLoaderData} from "react-router-dom";

export default () => {
  const { listener } = useLoaderData()

  return (
    <h1>{listener.name}</h1>
  )
}
