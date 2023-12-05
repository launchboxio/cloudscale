import {redirect} from "react-router-dom";
import axios from "axios";

const createTargetGroup = async({ request, params }) => {
  const formData = await request.formData();
  const payload = Object.fromEntries(formData)
  const res = await axios.post("/target_groups", payload)
  return redirect(`/target_groups/${res.data.target_group.id}`)
}

const addAttachment = async({request, params}) => {
  const { targetGroupId } = params
  const formData = await request.formData();
  const payload = Object.fromEntries(formData)
  payload.port = Number(payload.port)
  const res = await axios.post(`/target_groups/${targetGroupId}/attachments`, payload)
  return redirect(`/target_groups/${targetGroupId}`)
}

export default {
  createTargetGroup,
  addAttachment
}
