import type {Actions} from "../../../.svelte-kit/types/src/routes/(app)/settings/$types";
import {staticGoEP} from "$lib/models/enums/endpoints.enum";

export const actions : Actions = {
    uploadChatImg: async ({fetch, request}) => {
        let fd = await request.formData()
        console.log(fd)
        console.log(fd.get("fileName"))
        let result:boolean
        const resProm = await fetch(staticGoEP.CREATE_CHAT_IMG, {
            method: "POST",
            body: fd
        })

        result = resProm.ok;

        console.log(resProm.json())

        return {
            result
        }
    }
}