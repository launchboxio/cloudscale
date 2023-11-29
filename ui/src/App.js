import logo from './logo.svg';
import './App.css';
import Listener from "./widgets/listener";
import {Alignment, Button, Navbar} from "@blueprintjs/core";

import '@blueprintjs/core/lib/css/blueprint.css'
function App() {
  return (
    <>
      <Navbar>
        <Navbar.Group align={Alignment.LEFT}>
          <Navbar.Heading>Blueprint</Navbar.Heading>
          <Navbar.Divider />
          <Button className="bp5-minimal" icon="arrow-bottom-right" text="Listeners" />
          <Button className="bp5-minimal" icon="key" text="Certificates" />
          <Button className="bp5-minimal" icon="arrow-top-right" text="Target Groups" />

        </Navbar.Group>
      </Navbar>
      <Listener />
    </>
  );
}

export default App;
