<script setup>

</script>

<template>
  <main class="font-Roboto">
    <div class="bg-slate-400 text-black text-4xl pt-5 pb-5 pl-5">
      <h1>🍮 MamuroEmail</h1>
    </div>

    <div class="search mt-3 pl-2 pr-2">
      <div class="head-search">
        <i class="relative mr-2 left-2 z-10 top-8">🔎</i>
        <input @keyup="searchTerm()" class="w-full border rounded px-2 py-2 pl-9" type="text" name="search" id="search"
          placeholder="Escribe aquí" />
      </div>
      <div class="body-search flex mt-3">
        <div class="emails-details w-1/2 mr-6">
          <table class="mr-6 border-spacing-2 w-full">
            <thead class="">
              <tr class="bg-stone-300">
                <th class="border-2 border-current text-left">Subject</th>
                <th class="border-2 border-current text-left">From</th>
                <th class="border-2 border-current text-left">To</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(item, index) in items" :key="item._id" @click="getBody(item._source)"
                class="hover:bg-color-table cursor-pointer" :class="index % 2 ? '' : 'bg-zinc-200'">
                <td class="border-x-2 border-current" :class="index >= items.length - 1 ? 'border-b-2' : ''">{{
                    item._source.Subject
                }}</td>
                <td class="border-x-2 border-current" :class="index >= items.length - 1 ? 'border-b-2' : ''">{{
                    item._source.From
                }}</td>
                <td class="border-x-2 border-current" :class="index >= items.length - 1 ? 'border-b-2' : ''">{{
                    item._source.To
                }}</td>
              </tr>

              <!--<tr v-for="i in items.length + 9" class="hover:bg-color-table cursor-pointer"
                :class="i % 2 ? 'bg-zinc-200' : ''">
                <td class="border-x-2 border-current" :class="i > items.length + 8 ? 'border-b-2' : ''">&nbsp;</td>
                <td class="border-x-2 border-current" :class="i > items.length + 8 ? 'border-b-2' : ''">&nbsp;</td>
                <td class="border-x-2 border-current" :class="i > items.length + 8 ? 'border-b-2' : ''">&nbsp;</td>
              </tr>-->

            </tbody>
          </table>
        </div>
        <div class="emails-body w-1/2">
          <h3 class="font-bold pb-4" id="subjectText">{{ subjectText }}</h3>
          <div id="bodyText">
            <pre v-html="bodyText"></pre>
          </div>
        </div>
      </div>
    </div>


  </main>
</template>


<script>
import axios from "axios";
export default {
  name: "App",
  data() {
    return {
      items: [],
      subjectText: "No hay resultados",
      bodyText: "Sin resultados",
    };
  },
  methods: {
    initialData(hits) {
      if (hits.length > 0) {
        this.items = hits;
        this.subjectText = hits[0]._source.Subject
        this.bodyText = hits[0]._source.body
      }
    },
    getBody(data) {
      this.subjectText = data.Subject
      this.bodyText = data.body
    },
    async searchTerm() {
      const { value } = event.target
      console.log(value)
      const hits = await this.getData(value);
      this.initialData(hits)
    },
    async getData(term = "") {
      let query = ""
      if (term) {
        query = `?term=${term}`
      }
      try {
        const res = await axios.get(`http://localhost:3000/api/emails${query}`);
        const { hits } = res.data.hits;
        return hits
      } catch (error) {
        console.log(error);
      }
    },
  },
  async created() {
    try {
      //const res = await axios.get(`http://localhost:3000/api/emails`);
      //const { hits } = res.data.hits;
      const hits = await this.getData();
      this.initialData(hits)

    } catch (error) {
      console.log(error);
    }
  },
};
</script>
