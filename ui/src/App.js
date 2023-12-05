import logo from './logo.svg';
import './App.css';
import Listener from "./widgets/listener";
import {Alignment, Button, Navbar} from "@blueprintjs/core";


import {useState} from "react";
import Certificates from "./widgets/certificates";
import TargetGroups from "./widgets/target_groups";
function App() {
  const [location, setLocation] = useState("listeners")
  return (
    <>

      <div>
        {location === 'listeners' && <Listener />}
        {location === 'certificates' && <Certificates />}
        {location === 'target-groups' && <TargetGroups />}
      </div>
      {/*<Listener />*/}
    </>
  );
}

export default App;
