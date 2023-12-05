import {redirect} from "react-router-dom";
import axios from "axios";

const createListener = async({ request, params }) => {
  const formData = await request.formData();
  const payload = Object.fromEntries(formData)
  payload.port = Number(payload.port)
  const res = await axios.post("/listeners", payload)
  return redirect(`/listeners/${res.data.listener.id}`)
}

export default {
  createListener
}
