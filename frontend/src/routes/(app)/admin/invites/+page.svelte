<script lang="ts">
    import {authEP, userEP} from "$lib/models/enums/endpoints.enum";
    import AdminNav from "$lib/components/AdminNav.svelte";
    import {NewUserModel} from "$lib/models/classes/newUser.model";
    import ContentCopy from "svelte-material-icons/ContentCopy.svelte";
    import {onMount} from "svelte";
    import {get} from "$lib/utils/genericFetch";
    import FriendForm from "$lib/components/forms/FriendForm.svelte";
    import {newUserStore} from "$lib/stores/newUser.store";
    import type {INewUser} from "$lib/models/interfaces/inewUser.interface";

    onMount(async () => {
        $newUserStore = await get(userEP.REGISTERED)
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

</script>

<div class="h-full flex flex-col">
    <AdminNav page="friends"/>


    <h2 class="text-center">Invite your friends</h2>
    <div class="p-3">
        <FriendForm />
    </div>

    <div class="flex-1 overflow-auto">
        <div class="rounded max-h-1">
            {#each $newUserStore as user}
                <div class="p-3">
                    <div class="p-3 border-2 border-secondary rounded text-center">
                        <h3>{user.name}</h3>
                        <p>Send the magic link to {user.name} so they can log in</p>
                        <div class="flex justify-center space-x-2 p-3">
                            <button onclick={(e) => copyLink(user, e)}>
                                <span class="flex"><ContentCopy size="1.4em"/> Copy</span>
                            </button>
                        </div>
                    </div>
                </div>
            {/each}
        </div>
    </div>
</div>