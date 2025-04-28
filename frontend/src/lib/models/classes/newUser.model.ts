import type {INewUser} from "$lib/models/interfaces/inewUser.interface";

export class NewUserModel implements INewUser{
    public id!:      string;
    public name!:    string;
    public slug!:    string;
    public token!:   string;


    constructor(name: string) {
        this.name = name;
    }

    static deserialize(input: INewUser): NewUserModel {
        return {
            id: input.id,
            name: input.name,
            slug: input.slug,
            token: input.token
        } as NewUserModel
    }
}