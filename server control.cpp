#include "Bot1/Core/String1.h"

namespace Bot1::String1 {

UT_LINK(CoreCharString);
UT_LINK(CoreCharString1);
UT_LINK(CoreMSpace);
UT_LINK(CoreDateTime);

BOT1_SUBSYSTEM_INIT(UnitTest)
{
	testing::InitGoogleTest(BOT1_SUBSYSTEM_ARGC, BOT1_SUBSYSTEM_ARGV);
}

BOT1_SUBSYSTEM_UNEXPECTED_SHUTDOWN(UnitTest) {}

BOT1_SUBSYSTEM_DESTROY(UnitTest) {}

bool unit_test_all()
{
	return RUN_ALL_TESTS() == 0;
}

}
