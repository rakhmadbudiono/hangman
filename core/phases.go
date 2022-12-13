package core

type phase struct {
	desc string
	run  func(g *gallows)
}

var phases = []phase{
	{desc: "First, we draw a head", run: drawHead},
	{desc: "Now we draw a body.", run: drawBody},
	{desc: "Next we draw an arm.", run: drawRightArm},
	{desc: "this time it's the other arm.", run: drawLeftArm},
	{desc: "Now, let's draw the right leg.", run: drawRightLeg},
	{desc: "This time we draw the left leg.", run: drawLeftLeg},
	{desc: "Now we put up a hand.", run: drawLeftHand},
	{desc: "Next the other hand.", run: drawRightHand},
	{desc: "Now we draw one foot", run: drawLeftFoot},
	{desc: "Here's the other foot -- you're hung!!", run: drawRightFoot},
}
