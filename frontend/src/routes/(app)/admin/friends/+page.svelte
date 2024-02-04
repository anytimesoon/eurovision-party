<script lang="ts">
    import { enhance } from '$app/forms';
    import {authEP} from "$lib/models/enums/endpoints.enum";
    import type {PageData, ActionData} from "./$types";
    import AdminNav from "$lib/components/AdminNav.svelte";
    import type {NewUserModel} from "$lib/models/classes/user.model";

    export let data:PageData
    export let form:ActionData

    let users = data.users


    const updateUsers = (form) => {
        if(form != null) {
            users = [...users, form.user]
        }
    }

    const copyLink = (user:NewUserModel, e:Event) => {
        navigator.clipboard.writeText(authEP.SVELTE_LOGIN + user.token + "/" + user.id)
        let button = e.target as HTMLButtonElement
        button.innerText = "Copied!"
        setTimeout(function(){
            button.innerHTML = "<i class='fa-regular fa-copy'></i> Copy"
        }, 1000);
    }

    $: updateUsers(form)
</script>

<div class="h-full flex flex-col">
    <AdminNav page="friends"/>


    <h2 class="text-center">Invite your friends</h2>
    <div class="p-3">

        <form method="POST" action="?/register" use:enhance >
            <div class="w-fit mx-auto flex justify-center">
                <input class="mr-3" id="new-user-name" type="text" name="name" placeholder="Name"/>
                <button><i class="fa-regular fa-floppy-disk"></i></button>
            </div>

        </form>
    </div>

    <div class="flex-1 overflow-auto">
        <div class="rounded max-h-1">
            {#each users as user}
                <div class="p-3">
                    <div class="p-3 border-2 border-secondary rounded text-center">
                        <h3>{user.name}</h3>
                        <p>Send the magic link to {user.name} so they can log in</p>
                        <div class="flex justify-center space-x-2 p-3">
                            <button on:click={(e) => copyLink(user, e)}><i class="fa-regular fa-copy"></i> Copy</button>
    <!--                        <button ><i class="fa-solid fa-eye"></i> Show</button>-->
                        </div>
                    </div>
                </div>
            {/each}
        </div>
    </div>
</div>