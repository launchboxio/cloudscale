import axios from "axios";

const getTargetGroups = async () => {
  const res = await axios.get("/target_groups")
  return res.data
}

const getTargetGroup = async({ params }) => {
  const res = await axios.get(`/target_groups/${params.targetGroupId}`)
  return res.data
}

export default {
  getTargetGroups,
  getTargetGroup
}
