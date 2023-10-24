// import type {Actions} from "./$types";
// import {ResponseModel} from "$lib/models/classes/response.model";
// import type {VoteSingleModel} from "$lib/models/classes/voteSingle.model";
// import {voteGoEP} from "$lib/models/enums/endpoints.enum";
// import type {VoteModel} from "$lib/models/classes/vote.model";

// export const actions : Actions = {
//     vote: async ({fetch, request}) => {
//         const fd = await request.formData()
//         const vote:VoteSingleModel = Object.fromEntries([...fd]) as VoteSingleModel;
//
//         const resProm = await fetch(voteGoEP.UPDATE, {
//             method: "PUT",
//             body: JSON.stringify(vote,(key, value) => {
//                 if (key == "score") {
//                     return parseInt(value)
//                 }
//                 return value
//                 })
//         })
//
//         const res:ResponseModel<VoteModel> = await resProm.json()
//
//         return {
//             success: true,
//             vote: res.body,
//             error: res.error
//         }
//     }
// }