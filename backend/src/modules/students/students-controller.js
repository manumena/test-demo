const asyncHandler = require("express-async-handler");
const { getAllStudents, addNewStudent, getStudentDetail, setStudentStatus, updateStudent } = require("./students-service");

const handleGetAllStudents = asyncHandler(async (req, res) => {
    //write your code
    console.log("Fetching all students");
    const students = await getAllStudents(req.body);
    res.status(200).json(students);

});

const handleAddStudent = asyncHandler(async (req, res) => {
    //write your code
    console.log("Adding new student");
    const result = await addNewStudent(req.body);
    res.status(201).json(result);

});

const handleUpdateStudent = asyncHandler(async (req, res) => {
    //write your code
    console.log("Updating student");
    const result = await updateStudent({ ...req.body, id: req.params.id });
    res.status(200).json(result);
});

const handleGetStudentDetail = asyncHandler(async (req, res) => {
    //write your code
    console.log("Fetching student detail");
    const student = await getStudentDetail(req.params.id);
    res.status(200).json(student);
});

const handleStudentStatus = asyncHandler(async (req, res) => {
    //write your code
    console.log("Setting student status");
});

module.exports = {
    handleGetAllStudents,
    handleGetStudentDetail,
    handleAddStudent,
    handleStudentStatus,
    handleUpdateStudent,
};
