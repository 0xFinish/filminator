import axios from "axios";

const link = "http://localhost:8080"

export async function getQuestion(req) {
    console.log("request in getQuestion")
    console.log(req)
    const response = await axios.post(link + "/getRestrictions", req) 
    console.log(response.data)
    return response.data.question
}
