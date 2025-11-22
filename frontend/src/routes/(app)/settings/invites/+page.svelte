<script lang="ts">
    import {authEP, userEP} from "$lib/models/enums/endpoints.enum";
    import SettingsNav from "$lib/components/SettingsNav.svelte";
    import {NewUserModel} from "$lib/models/classes/newUser.model";
    import ContentCopy from "svelte-material-icons/ContentCopy.svelte";
    import {onMount} from "svelte";
    import {get, putWithError} from "$lib/utils/genericFetch";
    import FriendForm from "$lib/components/forms/FriendForm.svelte";
    import {newUserStore} from "$lib/stores/newUser.store";
    import {errorStore} from "$lib/stores/error.store";
    import type {INewUser} from "$lib/models/interfaces/inewUser.interface";
    import { userStore, currentUser } from "$lib/stores/user.store";
    import {flip} from "svelte/animate";
    import {fade} from "svelte/transition";

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

    const banUser = async (user: NewUserModel) => {
        const result = await putWithError<NewUserModel, NewUserModel>(userEP.BAN, user)
        console.log(result)
        if (user.id === result.body.id) {
            $newUserStore = $newUserStore.filter(u => u.id !== user.id)
        } else {
            $errorStore = result.error
        }
    }

    function getUserName(user: NewUserModel): string {
        const createdBy = $userStore.get(user.createdBy)
        return createdBy ? createdBy.name : ""
    }
</script>

<div class="h-full flex flex-col">
    <SettingsNav page="invites"/>

    <h2 class="text-center">Invite your friends</h2>
    <div class="p-3">

        <details class="overflow-hidden [&[open]_span::before]:rotate-90 [&[open]_span::before]:transition-transform [&[open]_span::before]:duration-200 [&[open]_span::before]:ease-out [&[open]~div.content]:max-h-96 [&[open]~div.content]:border-secondary [&[open]~div.content]:transition-all [&[open]~div.content]:duration-300 [&[open]~div.content]:ease-out [&[open]~div.content]:delay-0">
            <summary class="block [&::-webkit-details-marker]:hidden">
                <span class="relative flex items-center align-middle pl-4 bg-canvas-secondary h-16 cursor-pointer before:content-['►'] before:text-base before:flex before:items-center before:mr-2 before:transition-transform before:duration-200 before:delay-300 before:ease-out">
                    How does it work?
                </span>
            </summary>
        </details>
        <div class="content max-h-0 overflow-hidden border-x border-b border-secondary px-2.5 py-0 transition-all duration-500 ease-out delay-0">
            <ol class="list-decimal list-inside mb-2">
                <li>Put your friend's name in to create a magic ✨link</li>
                <li>Copy the ✨link and send it to them</li>
                <li>They can then log in using the ✨link and join the fun!</li>
            </ol>
            <p class="text-sm py-2">NOTE: the ✨link is unique to each person and can not be shared by multiple people.</p>
        </div>

    </div>

    <div class="p-3">
        <FriendForm />
    </div>

    <div class="flex-1 overflow-auto">
        <div class="rounded max-h-1">
            {#each $newUserStore as user (user.id)}
                <div class="p-3" animate:flip={{duration: 300}} in:fade>
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
                                <button onclick={() => banUser(user)}>
                                    <span class="flex">Ban</span>
                                </button>
                            </div>
                        </div>
                    </div>
                </div>
            {/each}
        </div>
    </div>
</div>