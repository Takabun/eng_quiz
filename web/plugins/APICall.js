import secrets from "../_secret/aws_info"
import AWS from "aws-sdk"
import {uuid} from "vue-uuid";
import "array-foreach-async";

const CreatenewOne = async (newTags, question, answer, questionImages, answerImages) => {
  for (let i = 0; i < newTags.length ; i ++ ) {
    const res = await axios.post(`api/tag`, { Name: newTags[i] })
    // this.question.Tags.push(res.data)
  };
  const questionresponse = await axios.post(`api/question`,question)
  const questionimageresponse = await onCreateQuestionImage(questionresponse.data.ID, questionImages)
  const answerresponse = await axios.post(`api/answer`, {
    Text: answer.Text,
    QuestionId: questionresponse.data.ID
  })
  const answerimageresponse = await onCreateAnswerImage(answerresponse.data.ID,  answerImages)
},  


const onCreateQuestionImage = async (qid, qImages) => {
  AWS.config.update({accessKeyId: secrets.accessKeyId, secretAccessKey: secrets.secretAccessKey});
  const bucket = new AWS.S3({params: {Bucket: secrets.Bucket}});
  if (this.questionImages[0]) {
    const params = {Key: uuid.v4() + "-" + qImages[0].name, ContentType: qImages[0].type, Body: qImages[0]}; 
    const uploadPromise = bucket.putObject(params).promise();
    uploadPromise
      .then(function(data) {
        console.log("uploaded!", data.$response);  
        axios.post(`api/questionimage`, {
          QuestionId: qid,
          Name:  data.$response.request.params.Body.name,
          Url: "https://quiztest-kt.s3-ap-northeast-1.amazonaws.com/" + data.$response.request.params.Key
        }).then(function(res) {
          console.log("DBへのQuestionImage保存まで完了", res)
        })
      })
      .catch(function(err) {
        console.error("error!", err, err.stack);
      });;
  }
},

const onCreateAnswerImage = async(aid, aImages) => {
  AWS.config.update({accessKeyId: secrets.accessKeyId, secretAccessKey: secrets.secretAccessKey});
  const bucket = new AWS.S3({params: {Bucket: secrets.Bucket}});
  if (this.answerImages[0]) {
    const params = {Key: uuid.v4() + "-" + aImages[0].name, ContentType: aImages[0].type, Body: aImages[0]};
    const uploadPromise = bucket.putObject(params).promise();
    uploadPromise
      .then(function(data) {
        console.log("uploaded!", data.$response); 
        axios.post(`api/answerimage`, {
          AnswerId: aid,
          Name:  data.$response.request.params.Body.name,
          Url: "https://quiztest-kt.s3-ap-northeast-1.amazonaws.com/" + data.$response.request.params.Key
        }).then(function(res) {
          console.log("DBへのAnswerImage保存まで完了", res)
        })

      })
      .catch(function(err) {
        console.error("error!", err, err.stack);
      });;
  }
}


export default ({}, inject) => {
  inject('CreatenewOne', CreatenewOne);
  inject('onCreateQuestionImage', onCreateQuestionImage);
  inject('onCreateAnswerImage', onCreateAnswerImage);
}