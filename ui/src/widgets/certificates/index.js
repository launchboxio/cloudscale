import React, {useEffect, useState} from "react";
import axios from "axios";
import {Button, Card, FormGroup, HTMLTable, InputGroup, Intent, Popover, TextArea} from "@blueprintjs/core";

const Certificates = () => {
  const [certificates, setCertificates] = useState([])
  const [newCertificate, setNewCertificate] = useState({})
  const [modalOpen, setModalOpen] = useState(false)

  useEffect(() => {
    axios.get("/certificates").then((res) => setCertificates(res.data.certificates))
  }, [])

  const handleSubmit = (event) => {
    event.preventDefault()

    axios.post("/certificates", newCertificate).then((res) => {
      setCertificates(certificates.concat([res.data.certificate]))
      setNewCertificate({})
      setModalOpen(false)
    })
  }

  return (
    <>
      <Card>
        <h4>Certificates</h4>
        <HTMLTable>
          <thead>
          <tr>
            <th>ID</th>
            <th>Name</th>
            <th>Domain</th>
            <th>Created At</th>
          </tr>
          </thead>
          <tbody>
          {certificates.map((item) => {
            return (
              <tr>
                <td>{item.id}</td>
                <td>{item.name}</td>
                <td>{item.domain}</td>
                <td>{item.created_at}</td>
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
                helperText="The friendly name of your certificate"
                label="Name"
                labelFor="name-input"
                labelInfo="(required)"
              >
                <InputGroup id="name-input" placeholder="Certificate name" onChange={(event) => {
                  setNewCertificate({
                    ...newCertificate,
                    name: event.target.value
                  })
                }}/>
              </FormGroup>
              <FormGroup
                helperText="The public cert for your certificate"
                label="Certificate"
                labelFor="cert-input"
              >
                <TextArea id="cert-input" onChange={(event) => {
                  setNewCertificate({
                    ...newCertificate,
                    cert: event.target.value
                  })
                }}/>
              </FormGroup>
              <FormGroup
                helperText="The private key"
                label="Private Key"
                labelFor="key-input"
                labelInfo="(required)"
              >
                <TextArea id="key-input" onChange={(event) => {
                  setNewCertificate({
                    ...newCertificate,
                    key: event.target.value
                  })
                }}/>
              </FormGroup>
              <Button intent={Intent.PRIMARY} text={"Submit"} type={"submit"} />
              <Button intent={Intent.NONE} text={"Cancel"} onClick={() => {
                setNewCertificate({})
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

export default Certificates
