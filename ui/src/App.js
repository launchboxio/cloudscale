import logo from './logo.svg';
import './App.css';
import Listener from "./widgets/listener";
import {Alignment, Button, Navbar} from "@blueprintjs/core";

import '@blueprintjs/core/lib/css/blueprint.css'
import {useState} from "react";
function App() {
  const [location, setLocation] = useState("listeners")
  return (
    <>
      <Navbar>
        <Navbar.Group align={Alignment.LEFT}>
          <Navbar.Heading>Blueprint</Navbar.Heading>
          <Navbar.Divider />
          <Button className="bp5-minimal" onClick={() => {
            setLocation("listeners")
          }} icon="arrow-bottom-right" text="Listeners" />
          <Button className="bp5-minimal" onClick={() => {
            setLocation("certificates")
          }} icon="key" text="Certificates" />
          <Button className="bp5-minimal" onClick={() => {
            setLocation("target-groups")
          }} icon="arrow-top-right" text="Target Groups" />

        </Navbar.Group>
      </Navbar>
      <div>
        {location === 'listeners' && <Listener />}
        {location === 'certificates' && <h1>Certificates</h1>}
        {location === 'target-groups' && <h1>Target Groups</h1>}
      </div>
      {/*<Listener />*/}
    </>
  );
}

export default App;
