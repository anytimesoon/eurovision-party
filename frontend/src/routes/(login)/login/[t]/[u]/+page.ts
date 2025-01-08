// @ts-ignore
import type {PageLoad} from ".$/types";

export const load:PageLoad =  async ({ params }) => {
    return {
        userId: params.u,
        loginToken: params.t
    }
}