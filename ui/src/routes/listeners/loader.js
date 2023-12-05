import axios from "axios";

const getListeners = async () => {
  const res = await axios.get("/listeners")
  return res.data
}

const getListener = async({ params }) => {
  const res = await axios.get(`/listeners/${params.listenerId}`)
  return res.data
}

export default {
  getListeners,
  getListener
}
