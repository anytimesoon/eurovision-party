<script lang="ts">
    import {authEP} from "$lib/models/enums/endpoints.enum.js";


</script>
    <table>
        <thead>
            <tr>
                <td>
                    Name
                </td>
                <td>
                    Email
                </td>
                <td>
                    Login
                </td>
            </tr>
        </thead>
        <tbody>
            {#each $registeredUserStore as user}
                <tr>
                <td>
                    <a href="{userFeEP.SINGLE_USER}{user.slug}">{user.name}</a>
                </td>
                <td>
                    {user.email}
                </td>
                <td>
                    <button on:click={navigator.clipboard.writeText(authEP.FE_LOGIN + user.token + "/" + user.id)}>Copy link</button>
                </td>
                </tr>
            {/each}
        </tbody>
    </table>
