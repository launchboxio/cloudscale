import * as React from "react"
import {Alignment, Button, Classes, Navbar} from "@blueprintjs/core";
import {Link, Outlet} from "react-router-dom";

const Root = () => {
  return (
    <>
      <Navbar>
        <Navbar.Group align={Alignment.LEFT}>
          <Navbar.Heading>Blueprint</Navbar.Heading>
          <Navbar.Divider />

          <Link to={"/listeners"}>
            <Button className={Classes.MINIMAL} icon="arrow-bottom-right" text="Listeners" />
          </Link>

          <Link to={"/certificates"}>
            <Button className={Classes.MINIMAL} icon="key" text="Certificates" />
          </Link>

          <Link to={"/target_groups"}>
            <Button className={Classes.MINIMAL} icon="arrow-top-right" text="Target Groups" />
          </Link>

        </Navbar.Group>
      </Navbar>
      <Outlet />
    </>
  )
}

export default Root
