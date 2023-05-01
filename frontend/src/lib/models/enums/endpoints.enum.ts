const apiBase = "http://localhost:8080/"
const feBase = "http://localhost:5173/"
const restricted = "restricted/api/"
const countryURL = apiBase + restricted + "country/"
const userURL = apiBase + restricted + "user/"
const voteURL = apiBase + restricted + "vote/"
//TODO: tidy up endpoints

// Backend endpoints
export const chatEP = "ws://localhost:8080/chat/connect";

export const countryEP = {
    ALL: countryURL,
    UPDATE: countryURL,
    FIND_ONE: countryURL,
    PARTICIPATING: countryURL + "participating"
}

export const staticEP = {
    IMG: apiBase
}

export const userEP = {
    ALL: userURL,
    UPDATE: userURL,
    FIND_ONE: userURL,
    REMOVE: userURL,
    REGISTERED: userURL + "registered"
}

export const voteEP = {
    CREATE: voteURL,
    UPDATE: voteURL,
    BY_COUNTRY: voteURL
}

export const authEP = {
    LOGIN: apiBase + "login",
    REGISTER: userURL + "register"
}

// Frontend enpoints

export const userFeEP = {
    SINGLE_USER: feBase + "person"
}