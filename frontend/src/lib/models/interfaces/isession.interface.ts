import {UserModel} from "$lib/models/classes/user.model";
import {CookieOpts} from "$lib/models/classes/session.model";

export interface ISession {
	name: 	string;
	token:  string;
	opts: 	CookieOpts;
	user: 	UserModel;
	botId: 	string;
}