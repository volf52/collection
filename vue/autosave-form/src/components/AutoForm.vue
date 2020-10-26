<template>
  <div>
    <div class="notification is-success" v-if="state === 'synced'">
      Form is synced with Firestore
    </div>
    <div class="notification is-link" v-else-if="state === 'modified'">
      From data changed, will sync with Firebase
    </div>
    <div class="notification is-warning" v-else-if="state === 'revoked'">
      From data and Firebase revoked to original data
    </div>
    <div class="notification is-danger" v-else-if="state === 'error'">
      Failed to save to Firestore. {{ errorMessage }}
    </div>
    <div class="notification is-info" v-else-if="state === 'loading'">
      Loading...
    </div>

    <hr />

    <div class="columns">
      <div class="column">
        <h3>Original Data</h3>
        {{ originalData }}
      </div>
      <div class="column">
        <h3>Vue Form Data</h3>
        <br />
        {{ formData }}
      </div>
      <div class="column">
        <h3>Firebase Data</h3>
        <br />
        {{ firebaseData }}
      </div>
    </div>

    <hr />

    <form @submit.prevent="updateFirebase" @input="fieldUpdate">
      <div class="field is-horizontal">
        <div class="field-label is-normal">
          <label for="contact" class="label">Contact</label>
        </div>
        <div class="field-body">
          <div class="field">
            <p class="control is-expanded has-icons-left">
              <input
                class="input"
                type="text"
                name="name"
                v-model="formData.name"
              />
              <span class="icon is-small is-left"
                ><i class="fas fa-user"></i
              ></span>
            </p>
          </div>

          <div class="field">
            <p class="control is-expanded has-icons-left has-icons-right">
              <input
                class="input"
                type="email"
                name="email"
                v-model="formData.email"
              />
              <span class="icon is-small is-left"
                ><i class="fas fa-envelope"></i
              ></span>
            </p>
          </div>
        </div>
      </div>

      <div class="field is-horizontal">
        <div class="field-label is-normal">
          <label class="label">Phone</label>
        </div>
        <div class="field-body">
          <div class="field is-expanded">
            <div class="field has-addons">
              <p class="control"><a class="button is-static">+90</a></p>
              <p class="control is-expanded">
                <input
                  class="input"
                  type="tel"
                  name="phone"
                  v-model="formData.phone"
                />
              </p>
            </div>
            <p class="help">Do not enter the first zero</p>
          </div>
        </div>
      </div>
      <hr />
      Vue Form State: {{ state === "" ? "waiting for changed" : state }}
      <hr />

      <button
        class="button is-success"
        type="submit"
        v-if="state === 'modified'"
      >
        Save Changes
      </button>
    </form>

    <br />
    <button class="button is-warning is-rounded" @click="revertToOriginal">
      <span class="icon"><i class="fas fa-undo"></i></span>
      <span>Revoke to Original data</span>
    </button>
  </div>
</template>

<script>
import { db } from "../firebase";
import { debounce } from "debounce";

const docPath = "contacts/jeff";

export default {
  name: "AutoForm",

  data() {
    return {
      state: "loading", // synced, modified, revoked, error
      firebaseData: null,
      formData: {},
      errMsg: "",

      originalData: null
    };
  },

  firestore() {
    return {
      firebaseData: db.doc(docPath)
    };
  },

  methods: {
    async updateFirebase() {
      try {
        await db.doc(docPath).set(this.formData);
        this.state = "synced";
      } catch (err) {
        console.error(err);
        this.errMsg = JSON.stringify(err);
        this.state = "error";
      }
    },

    fieldUpdate() {
      this.state = "modified";
      this.debouncedUpdate();
    },

    debouncedUpdate: debounce(function() {
      this.updateFirebase();
    }, 1500),

    revertToOriginal() {
      this.state = "revoked";
      this.formData = { ...this.originalData };
    }
  },

  created: async function() {
    const docRef = db.doc(docPath);
    let data = (await docRef.get()).data();

    if (!data) {
      data = { name: "", phone: "", email: "" };
      docRef.set(data);
    }

    this.formData = data;
    this.originalData = { ...data };
    this.state = "synced";
  }
};
</script>

<style scoped>
h1,
h2,
h3 {
  font-weight: 600;
}

p.help {
  text-align: left;
}
</style>