import type { ISession } from '../interfaces/isession.interface';
import type {UserModel} from "$lib/models/classes/user.model";
import type {sameSite} from "$lib/models/enums/cookie.enum";

export class SessionModel implements ISession {
    public name!:   string;
    public token!:  string;
    public opts!:   CookieOpts;
    public user!:   UserModel;
    public botId!:  string;
}

export class CookieOpts {
    public path: string;
    public maxAge: number;
    public secure: boolean;
    public HttpOnly: boolean;
    public SameSite: sameSite;
    public Domain: string;
}
