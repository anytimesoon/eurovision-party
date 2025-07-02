<script lang="ts">
    import {authEP, userEP} from "$lib/models/enums/endpoints.enum";
    import SettingsNav from "$lib/components/SettingsNav.svelte";
    import {NewUserModel} from "$lib/models/classes/newUser.model";
    import ContentCopy from "svelte-material-icons/ContentCopy.svelte";
    import {onMount} from "svelte";
    import {get} from "$lib/utils/genericFetch";
    import FriendForm from "$lib/components/forms/FriendForm.svelte";
    import {newUserStore} from "$lib/stores/newUser.store";
    import type {INewUser} from "$lib/models/interfaces/inewUser.interface";
    import { userStore, currentUser } from "$lib/stores/user.store";

    onMount(async () => {
        $newUserStore = await get(userEP.REGISTERED+$currentUser.id)
            .then((res: Array<INewUser>) =>
                res.map((user: INewUser) => NewUserModel.deserialize(user))
            )
    })

    const copyLink = (user:NewUserModel, e:Event) => {
        navigator.clipboard.writeText(authEP.SVELTE_LOGIN + user.token + "/" + user.id)
        let button = e.target as HTMLButtonElement
        const content = button.innerHTML
        button.innerText = "Copied!"
        setTimeout(function(){
            button.innerHTML = content
        }, 1000);
    }

    function getUserName(user: NewUserModel): string {
        const createdBy = $userStore.get(user.createdBy)
        return createdBy ? createdBy.name : ""
    }
</script>

<div class="h-full flex flex-col">
    <SettingsNav page="invites"/>

    <h2 class="text-center">Invite your friends</h2>
    <div>
        <details>
            <summary class="px-5 py-3 text-lg cursor-pointer text-center">How does it work?</summary>
            <div class="px-5 py-3 border border-gray-300">
                <ol class="list-decimal list-inside mb-2">
                    <li>Put your friend's name in to create a magic ✨link</li>
                    <li>Copy the ✨link and send it to them</li>
                    <li>They can then log in using the ✨link and join the fun!</li>
                </ol>
                <p class="text-sm">NOTE: the ✨link is unique to each person and can not be shared by multiple people.</p>
            </div>
        </details>
    </div>

    <div class="p-3">
        <FriendForm />
    </div>

    <div class="flex-1 overflow-auto">
        <div class="rounded max-h-1">
            {#each $newUserStore as user}
                <div class="p-3">
                    <div class="p-3 border-2 border-secondary rounded">
                        <div class="grid grid-cols-2 grid-gap-2">
                            <div>
                                <h3>{user.name}</h3>
                                <p class="text-sm">Invited By: {getUserName(user)}</p>
                            </div>
                            <div class="text-right">
                                <button onclick={(e) => copyLink(user, e)}>
                                    <span class="flex"><ContentCopy size="1.4em"/> Copy ✨</span>
                                </button>
                            </div>
                        </div>
                    </div>
                </div>
            {/each}
        </div>
    </div>
</div>