<template>

  <div id="survey-management">
    <h1>School List</h1>
    <button @click="fetchSchools">Get Schools</button>
    <div v-if="error">{{ error }}</div>
    <ul v-else>
      <li v-for="school in schools" :key="school.Code">
        <strong>{{ school.Name }}</strong> - {{ school.Teacher }} teaches {{ school.Subject }}
      </li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  data() {
    return {
      schools: [],
      error: null
    };
  },
  methods: {
    fetchSchools() {
      axios.get('http://localhost:3000/schools')
        .then(response => {
          this.schools = response.data.Data;
        })
        .catch(error => {
          this.error = 'Failed to load schools: ' + error.message;
        });
    }
  }
}
</script>

<style>

#survey-management {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
